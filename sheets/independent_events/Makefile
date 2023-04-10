# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make independent_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

independent_events.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > independent_events.tex

independent_events.pdf: ../../*.tex ../../trademark.pdf *.tex independent_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode independent_events.tex
	../../latexrun independent_events.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_events.pdf
	open independent_events.pdf

o:
	make open

terms: independent_events.tex
	bbk_sheet -mode ts independent_events.tex

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
	rm -f independent_events.*
	make reset
	bbk_sheet -mode mk > Makefile