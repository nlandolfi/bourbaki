# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make marginal_distributions.pdf
	make graph.pdf

marginal_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > marginal_distributions.tex

marginal_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_distributions.tex
	pdflatex --file-line-error -interaction=nonstopmode marginal_distributions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: marginal_distributions.pdf
	open marginal_distributions.pdf

o:
	make open

defs: marginal_distributions.tex
	defs marginal_distributions.tex

spell:
	aspell -c marginal_distributions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm marginal_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile