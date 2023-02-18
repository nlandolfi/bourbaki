# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cumulative_distributions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

cumulative_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cumulative_distributions.tex

cumulative_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cumulative_distributions.tex
	../../latexrun cumulative_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distributions.pdf
	open cumulative_distributions.pdf

o:
	make open

terms: cumulative_distributions.tex
	bbk_sheet -mode ts cumulative_distributions.tex

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
	rm -f cumulative_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile