# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make number_of_disjoint_unions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

number_of_disjoint_unions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > number_of_disjoint_unions.tex

number_of_disjoint_unions.pdf: ../../*.tex ../../trademark.pdf *.tex number_of_disjoint_unions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode number_of_disjoint_unions.tex
	../../latexrun number_of_disjoint_unions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: number_of_disjoint_unions.pdf
	open number_of_disjoint_unions.pdf

o:
	make open

terms: number_of_disjoint_unions.tex
	bbk sheet -mode ts number_of_disjoint_unions.tex

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
	rm -f number_of_disjoint_unions.*
	make reset
	bbk sheet -mode mk > Makefile