# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make parametric_distribution_network_families.pdf

parametric_distribution_network_families.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > parametric_distribution_network_families.tex

parametric_distribution_network_families.pdf: ../../*.tex ../../trademark.pdf *.tex parametric_distribution_network_families.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode parametric_distribution_network_families.tex
	../../latexrun parametric_distribution_network_families.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: parametric_distribution_network_families.pdf
	open parametric_distribution_network_families.pdf

o:
	make open

terms: parametric_distribution_network_families.tex
	bbk_sheet -mode ts parametric_distribution_network_families.tex

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
	rm -f parametric_distribution_network_families.*
	make reset
	bbk_sheet -mode mk > Makefile