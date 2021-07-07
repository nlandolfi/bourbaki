# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make marginal_distributions.pdf

marginal_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > marginal_distributions.tex

marginal_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode marginal_distributions.tex
	../../latexrun marginal_distributions.tex
	make terms

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

terms: marginal_distributions.tex
	bbk_sheet -mode ts marginal_distributions.tex

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
	rm -f marginal_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile