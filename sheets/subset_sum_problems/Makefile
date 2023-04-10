# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subset_sum_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

subset_sum_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > subset_sum_problems.tex

subset_sum_problems.pdf: ../../*.tex ../../trademark.pdf *.tex subset_sum_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subset_sum_problems.tex
	../../latexrun subset_sum_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subset_sum_problems.pdf
	open subset_sum_problems.pdf

o:
	make open

terms: subset_sum_problems.tex
	bbk sheet -mode ts subset_sum_problems.tex

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
	rm -f subset_sum_problems.*
	make reset
	bbk sheet -mode mk > Makefile