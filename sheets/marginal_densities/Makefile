# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make marginal_densities.pdf
	make graph.pdf

marginal_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > marginal_densities.tex

marginal_densities.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_densities.tex
	# pdflatex --file-line-error -interaction=nonstopmode marginal_densities.tex
	../../latexrun marginal_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: marginal_densities.pdf
	open marginal_densities.pdf

o:
	make open

terms: marginal_densities.tex
	bbk_sheet -mode ts marginal_densities.tex

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
	rm -f marginal_densities.*
	make reset
	bbk_sheet -mode mk > Makefile