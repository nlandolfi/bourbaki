# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_differential_entropy.pdf

normal_differential_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > normal_differential_entropy.tex

normal_differential_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex normal_differential_entropy.tex
	# pdflatex --file-line-error -interaction=nonstopmode normal_differential_entropy.tex
	../../latexrun normal_differential_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_differential_entropy.pdf
	open normal_differential_entropy.pdf

o:
	make open

terms: normal_differential_entropy.tex
	bbk_sheet -mode ts normal_differential_entropy.tex

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
	rm -f normal_differential_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile