# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tree_approximators_of_a_normal.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

tree_approximators_of_a_normal.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > tree_approximators_of_a_normal.tex

tree_approximators_of_a_normal.pdf: ../../*.tex ../../trademark.pdf *.tex tree_approximators_of_a_normal.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode tree_approximators_of_a_normal.tex
	../../latexrun tree_approximators_of_a_normal.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_approximators_of_a_normal.pdf
	open tree_approximators_of_a_normal.pdf

o:
	make open

terms: tree_approximators_of_a_normal.tex
	bbk sheet -mode ts tree_approximators_of_a_normal.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree_approximators_of_a_normal.*
	make reset
	bbk sheet -mode mk > Makefile