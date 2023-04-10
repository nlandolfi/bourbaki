# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make marginal_densities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

marginal_densities.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > marginal_densities.tex

marginal_densities.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_densities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode marginal_densities.tex
	../../latexrun marginal_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

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