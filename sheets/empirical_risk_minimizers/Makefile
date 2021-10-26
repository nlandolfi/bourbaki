# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empirical_risk_minimizers.pdf

empirical_risk_minimizers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empirical_risk_minimizers.tex

empirical_risk_minimizers.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_risk_minimizers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empirical_risk_minimizers.tex
	../../latexrun empirical_risk_minimizers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_risk_minimizers.pdf
	open empirical_risk_minimizers.pdf

o:
	make open

terms: empirical_risk_minimizers.tex
	bbk_sheet -mode ts empirical_risk_minimizers.tex

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
	rm -f empirical_risk_minimizers.*
	make reset
	bbk_sheet -mode mk > Makefile