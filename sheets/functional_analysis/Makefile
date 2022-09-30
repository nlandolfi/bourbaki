# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make functional_analysis.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

functional_analysis.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > functional_analysis.tex

functional_analysis.pdf: ../../*.tex ../../trademark.pdf *.tex functional_analysis.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode functional_analysis.tex
	../../latexrun functional_analysis.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: functional_analysis.pdf
	open functional_analysis.pdf

o:
	make open

terms: functional_analysis.tex
	bbk_sheet -mode ts functional_analysis.tex

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
	rm -f functional_analysis.*
	make reset
	bbk_sheet -mode mk > Makefile