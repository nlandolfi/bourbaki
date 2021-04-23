# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sigma_algebra_independence.pdf
	make graph.pdf

sigma_algebra_independence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sigma_algebra_independence.tex

sigma_algebra_independence.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebra_independence.tex
	# pdflatex --file-line-error -interaction=nonstopmode sigma_algebra_independence.tex
	../../latexrun sigma_algebra_independence.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sigma_algebra_independence.pdf
	open sigma_algebra_independence.pdf

o:
	make open

terms: sigma_algebra_independence.tex
	bbk_sheet -mode ts sigma_algebra_independence.tex

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
	rm -f sigma_algebra_independence.*
	make reset
	bbk_sheet -mode mk > Makefile